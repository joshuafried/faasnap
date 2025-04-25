#!/bin/bash

set -e
set -x

wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.26/swagger-codegen-cli-2.4.26.jar
java -jar ~/swagger-codegen-cli-2.4.26.jar generate -i api/swagger.yaml -l python  -o python-client

