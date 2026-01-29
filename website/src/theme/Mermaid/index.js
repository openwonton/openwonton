import React, { useEffect, useRef, useState } from 'react';
import BrowserOnly from '@docusaurus/BrowserOnly';
import {
  MermaidContainerClassName,
  useMermaidConfig,
} from '@docusaurus/theme-mermaid/client';
import mermaid from 'mermaid';
import styles from './styles.module.css';

function MermaidDiagram({ value }) {
  const mermaidConfig = useMermaidConfig();
  const containerRef = useRef(null);
  const [svg, setSvg] = useState('');

  useEffect(() => {
    let cancelled = false;

    const renderDiagram = async () => {
      try {
        const init = mermaid.initialize ?? mermaid.mermaidAPI?.initialize;
        if (init) {
          init(mermaidConfig);
        }

        const render = mermaid.render ?? mermaid.mermaidAPI?.render;
        if (!render) {
          throw new Error('Mermaid renderer is unavailable.');
        }

        const id = `mermaid-svg-${Math.round(Math.random() * 10000000)}`;
        const result = await render(id, value);
        const nextSvg = typeof result === 'string' ? result : result.svg;

        if (cancelled) {
          return;
        }

        setSvg(nextSvg);

        if (typeof result === 'object' && result?.bindFunctions) {
          requestAnimationFrame(() => {
            if (!cancelled && containerRef.current) {
              result.bindFunctions(containerRef.current);
            }
          });
        }
      } catch (error) {
        if (!cancelled) {
          setSvg(`<pre>Mermaid render error: ${String(error)}</pre>`);
        }
      }
    };

    renderDiagram();

    return () => {
      cancelled = true;
    };
  }, [value, mermaidConfig]);

  return (
    <div
      ref={containerRef}
      className={`${MermaidContainerClassName} ${styles.container}`}
      // eslint-disable-next-line react/no-danger
      dangerouslySetInnerHTML={{ __html: svg }}
    />
  );
}

export default function Mermaid(props) {
  return <BrowserOnly>{() => <MermaidDiagram {...props} />}</BrowserOnly>;
}
