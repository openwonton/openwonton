// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

import React from 'react';
import MDXComponents from '@theme-original/MDXComponents';
import Admonition from '@theme/Admonition';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import CodeBlockConfig from '@site/src/components/CodeBlockConfig';
import EnterpriseAlert from '@site/src/components/EnterpriseAlert';
import Placement from '@site/src/components/Placement';

function createAdmonition(type, defaultTitle) {
  return function AdmonitionShim({ title, children, ...props }) {
    const resolvedTitle = title || defaultTitle;
    return (
      <Admonition type={type} title={resolvedTitle} {...props}>
        {children}
      </Admonition>
    );
  };
}

const Tip = createAdmonition('tip');
const Note = createAdmonition('note');
const Warning = createAdmonition('warning');
const Caution = createAdmonition('caution');
const Highlight = createAdmonition('info', 'Highlight');

function inferGroup(children) {
  const childArray = React.Children.toArray(children);
  for (const child of childArray) {
    if (React.isValidElement(child) && child.props && child.props.group) {
      return child.props.group;
    }
  }
  return undefined;
}

function TabsShim({ children, group, groupId, ...props }) {
  const childArray = React.Children.toArray(children).filter((child) =>
    React.isValidElement(child)
  );
  const resolvedGroup = groupId || group || inferGroup(childArray);
  const normalizedChildren = childArray.map((child, index) => {
    if (!React.isValidElement(child)) {
      return null;
    }

    if ('value' in child.props) {
      return child;
    }

    const { heading, group: tabGroup, label, value, ...rest } = child.props;
    const resolvedLabel = label || heading;
    const resolvedValue =
      value || heading || label || `tab-${index}`;

    return (
      <TabItem
        key={child.key ?? resolvedValue}
        label={resolvedLabel}
        value={resolvedValue}
        {...rest}
      />
    );
  });

  return (
    <Tabs {...props} groupId={resolvedGroup}>
      {normalizedChildren}
    </Tabs>
  );
}

function Tab({ children }) {
  return <>{children}</>;
}

export default {
  ...MDXComponents,
  Tip,
  Note,
  Warning,
  Caution,
  Highlight,
  Tabs: TabsShim,
  Tab,
  CodeBlockConfig,
  EnterpriseAlert,
  Placement,
};
