/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

import ApplicationSerializer from './application';

export default ApplicationSerializer.extend({
  serialize() {
    const json = ApplicationSerializer.prototype.serialize.apply(
      this,
      arguments
    );
    if (json instanceof Array) {
      json.forEach(serializeRole);
    } else {
      serializeRole(json);
    }
    return json;
  },
});

function serializeRole(role) {
  const policyIds = role.PolicyIDs || [];
  role.Policies = policyIds.map((policy) => {
    return { Name: policy };
  });
  delete role.PolicyIDs;
  return role;
}
