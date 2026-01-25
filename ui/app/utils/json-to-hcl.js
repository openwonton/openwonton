/**
 * Copyright (c) 2025 OpenWonton Authors.
 * SPDX-License-Identifier: MPL-2.0
 */

// Clean-room replacement; see CLEAN_ROOM_NOTES.md.
const IDENTIFIER_RE = /^[A-Za-z_][A-Za-z0-9_]*$/;
const BOOLEAN_RE = /^(true|false)$/i;
const NUMBER_RE = /^-?\d+(\.\d+)?([eE][+-]?\d+)?$/;

export default function jsonToHcl(flags) {
  if (!flags || typeof flags !== 'object') {
    return '';
  }

  const keys = Object.keys(flags);
  if (keys.length === 0) {
    return '';
  }

  keys.sort();

  let output = '';
  for (const key of keys) {
    output += `${formatKey(key)} = ${formatValue(flags[key])}\n`;
  }

  return output;
}

function formatKey(key) {
  const rawKey = String(key);
  if (IDENTIFIER_RE.test(rawKey)) {
    return rawKey;
  }

  return `"${escapeString(rawKey)}"`;
}

function formatValue(value) {
  if (value === null || value === undefined) {
    return '""';
  }

  if (typeof value === 'boolean' || typeof value === 'number') {
    return String(value);
  }

  const raw = String(value);
  const trimmed = raw.trim();

  if (!trimmed) {
    return '""';
  }

  const quote = trimmed[0];
  if (
    (quote === '"' || quote === "'") &&
    trimmed.length >= 2 &&
    trimmed[trimmed.length - 1] === quote
  ) {
    return trimmed;
  }

  if (BOOLEAN_RE.test(trimmed)) {
    return trimmed;
  }

  if (NUMBER_RE.test(trimmed)) {
    return trimmed;
  }

  if (
    (trimmed.startsWith('{') && trimmed.endsWith('}')) ||
    (trimmed.startsWith('[') && trimmed.endsWith(']'))
  ) {
    return trimmed;
  }

  return `"${escapeString(raw)}"`;
}

function escapeString(value) {
  return value
    .replace(/\\/g, '\\\\')
    .replace(/"/g, '\\"')
    .replace(/\n/g, '\\n')
    .replace(/\t/g, '\\t')
    .replace(/\r/g, '\\r');
}
