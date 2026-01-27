// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

import React from 'react';

const defaultCopy =
  'This feature is available in OpenWonton Enterprise.';

export default function EnterpriseAlert({ inline = false, children }) {
  if (inline) {
    return <span className="enterprise-alert-inline">Enterprise</span>;
  }

  return (
    <div className="enterprise-alert">
      <div className="enterprise-alert__title">Enterprise</div>
      <div className="enterprise-alert__body">
        {children || defaultCopy}
      </div>
    </div>
  );
}
