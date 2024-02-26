#!/bin/bash

TFVARS_FILE="thanksfulness-prd.tfvars"
TEMPLATE_FILE="thanksfulness-prd.tfvars.template"

if [ -f "$TFVARS_FILE" ]; then
    echo "$TFVARS_FILE file exists. Exiting..."
    exit 0
else
    if [ -f "$TEMPLATE_FILE" ]; then
        cp "$TEMPLATE_FILE" "$TFVARS_FILE"
        echo "Created $TFVARS_FILE file from $TEMPLATE_FILE"
    else
        echo "$TEMPLATE_FILE file not found. Unable to create $TFVARS_FILE file."
        exit 1
    fi
fi
