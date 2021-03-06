#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from gql.gql.client import OperationException
from gql.gql.reporter import FailedOperationException
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional, Dict
from time import perf_counter
from dataclasses_json import DataClassJsonMixin

from ..fragment.property import PropertyFragment, QUERY as PropertyFragmentQuery

QUERY: List[str] = PropertyFragmentQuery + ["""
query EquipmentTypeAndPropertiesQuery($id: ID!) {
  equipment: node(id: $id) {
    ... on Equipment {
      equipmentType {
        name
      }
      properties {
        ...PropertyFragment
      }
    }
  }
}

"""]

@dataclass
class EquipmentTypeAndPropertiesQuery(DataClassJsonMixin):
    @dataclass
    class EquipmentTypeAndPropertiesQueryData(DataClassJsonMixin):
        @dataclass
        class Node(DataClassJsonMixin):
            @dataclass
            class EquipmentType(DataClassJsonMixin):
                name: str

            @dataclass
            class Property(PropertyFragment):
                pass

            equipmentType: EquipmentType
            properties: List[Property]

        equipment: Optional[Node]

    data: EquipmentTypeAndPropertiesQueryData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str) -> Optional[EquipmentTypeAndPropertiesQueryData.Node]:
        # fmt: off
        variables: Dict[str, Any] = {"id": id}
        try:
            network_start = perf_counter()
            response_text = client.call(''.join(set(QUERY)), variables=variables)
            decode_start = perf_counter()
            res = cls.from_json(response_text).data
            decode_time = perf_counter() - decode_start
            network_time = decode_start - network_start
            client.reporter.log_successful_operation("EquipmentTypeAndPropertiesQuery", variables, network_time, decode_time)
            return res.equipment
        except OperationException as e:
            raise FailedOperationException(
                client.reporter,
                e.err_msg,
                e.err_id,
                "EquipmentTypeAndPropertiesQuery",
                variables,
            )
