# Copyright (c) 2004-present Facebook All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Extend base fluentd image
FROM frinx/fm-conductor-server:latest

# Copy config files
COPY config.properties /app/config/config.properties
COPY log4j-fluentd-appender.properties /app/config/log4j-fluentd-appender.properties
# Enable es6 compatibility for conductor's lambda task
RUN sed -i 's/java -jar/java -jar -Dnashorn.args=--language=es6/g' /app/startup.sh
