/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';

export default function useWasm() {
  const [isReady, setIsReady] = useState(false);

  useEffect(() => {
    const loadWasm = async () => {
      const script = document.createElement('script');
      
      script.src = '/wasm_exec.js';
      script.onload = async () => {
        const go = new (window as any).window.Go();
        const wasm = await WebAssembly.instantiateStreaming(
          fetch('/main.wasm'),
          go.importObject
        );
        go.run(wasm.instance);
        setIsReady(true);
      };
      document.body.appendChild(script);
    };

    loadWasm();
  }, []);

  return { isReady };
}
