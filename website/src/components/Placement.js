// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

import React from 'react';

function normalizeGroups(groups) {
  if (!Array.isArray(groups)) {
    return [];
  }

  if (groups.length === 0) {
    return [];
  }

  if (Array.isArray(groups[0])) {
    return groups;
  }

  return [groups];
}

export default function Placement({ groups }) {
  const normalized = normalizeGroups(groups);
  if (!normalized.length) {
    return null;
  }

  const label = normalized
    .map((path) => path.filter(Boolean).join(' → '))
    .join(' · ');

  return (
    <div className="placement">
      <span className="placement__label">Placement:</span>
      <span>{label}</span>
    </div>
  );
}
