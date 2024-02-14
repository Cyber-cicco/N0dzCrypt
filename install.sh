#!/bin/bash

rm -rf ~/go/bin/resources/ && cp -r resources/ ~/go/bin/resources/ && cd src && go install
