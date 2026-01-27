// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

import React from 'react';
import CodeBlock from '@theme/CodeBlock';

function getCodeProps(child) {
  if (!React.isValidElement(child)) {
    return null;
  }

  if (child.type === 'pre' && React.isValidElement(child.props.children)) {
    const codeChild = child.props.children;
    const className = codeChild.props.className || '';
    const language = className.replace(/^language-/, '');
    return {
      language,
      code: codeChild.props.children,
    };
  }

  const className = child.props.className || '';
  const language = className.replace(/^language-/, '');
  return {
    language,
    code: child.props.children,
  };
}

export default function CodeBlockConfig({ children, filename, hideClipboard }) {
  const child = React.Children.only(children);
  const codeProps = getCodeProps(child);

  if (!codeProps) {
    return children;
  }

  return (
    <CodeBlock
      language={codeProps.language}
      title={filename}
      showCopyButton={!hideClipboard}
    >
      {codeProps.code}
    </CodeBlock>
  );
}
