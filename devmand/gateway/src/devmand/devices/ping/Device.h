// Copyright (c) 2016-present, Facebook, Inc.
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree. An additional grant
// of patent rights can be found in the PATENTS file in the same directory.

#pragma once

#include <devmand/channels/ping/Channel.h>
#include <devmand/devices/Device.h>

namespace devmand {
namespace devices {
namespace ping {

class Device : public devices::Device {
 public:
  Device(
      Application& application,
      const Id& id,
      bool readonly_,
      const folly::IPAddress& ip_);
  Device() = delete;
  virtual ~Device() = default;
  Device(const Device&) = delete;
  Device& operator=(const Device&) = delete;
  Device(Device&&) = delete;
  Device& operator=(Device&&) = delete;

  static std::shared_ptr<devices::Device> createDevice(
      Application& app,
      const cartography::DeviceConfig& deviceConfig);

 public:
  std::shared_ptr<Datastore> getOperationalDatastore() override;

 protected:
  void setIntendedDatastore(const folly::dynamic& config) override {
    (void)config;
    LOG(ERROR) << "set config on unconfigurable device";
  }

 protected:
  channels::ping::Channel channel;
};

} // namespace ping
} // namespace devices
} // namespace devmand
