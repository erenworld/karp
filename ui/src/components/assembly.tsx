// useEffect(() => {
//   const loadWasm = async () => {
//     const go = new window.Go();
//     const result = await WebAssembly.instantiateStreaming(
//       fetch('/monkey.wasm'),
//       go.importObject
//     );
//     go.run(result.instance);
//   };

//   const script = document.createElement('script');
//   script.src = '/wasm_exec.js';
//   script.onload = loadWasm;
//   document.body.appendChild(script);
// }, []);
