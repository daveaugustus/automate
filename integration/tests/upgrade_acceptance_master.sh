#!/bin/bash

#shellcheck disable=SC2034
test_name="upgrade_acceptance_master"
test_upgrades=true
test_upgrade_inspec_profiles=()
test_diagnostics_filters="~iam-v2"
test_upgrade_begin_manifest="acceptance.json"
