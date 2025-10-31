/* eslint-disable @typescript-eslint/no-explicit-any */
// import CodeInputContent from './components/input-content';
// import ResultsContent from './components/result-content';

// function App() {
//   return (
//     <div className="">
//       <h1 className="text-[#ddd] text-xl">Karp interpreter</h1>
//       <p className="text-[#9d9d9d]">Small elegent interpter for fun</p>
//       <div className="py-4 text-gray-300 font-light">
//         <p>code example</p>
//         <div className="space-x-3.5">
//           <button>Hello world!</button>
//           <button>Fibonacci</button>
//           <button>Factorial</button>
//           <button>Addition</button>
//           <button>Map</button>
//           <button>Hashmap</button>
//           <button>Array indexing</button>
//         </div>
//       </div>
//       <div className="w-2/4 h-screen">
//         <button className="cursor-pointer bg-gray-200 border text-black border-black px-3 py-2 mr-3">
//           View on github
//         </button>
//         <button className="cursor-pointer bg-black border border-gray-300 text-gray-300 font-light px-3 py-2 mr-3">
//           Install binary
//         </button>
//         <button className="cursor-pointer bg-black border border-gray-300 text-gray-300 font-light px-3 py-2 mr-3">
//           Clear
//         </button>
//         <button className="cursor-pointer bg-black border border-gray-300 text-gray-300 font-light px-3 py-2 mr-3">
//           Share
//         </button>
//         <button className="cursor-pointer bg-black border border-gray-300 text-gray-300 font-light px-3 py-2 mr-3">
//           Run
//         </button>
//         <CodeInputContent input="" setInput={'dd'} />
//       </div>
//       <div>
//         <ResultsContent />
//       </div>
//     </div>
//   );
// }

// export default App;
import { useState } from 'react';
import useKarpWasm from './hooks/useWasm';

function App() {
  const { isReady } = useKarpWasm();
  const [input, setInput] = useState('');
  const [output, setOutput] = useState('');
  const [view, setView] = useState('output');

  const handleRun = async () => {
    if (!isReady) {
      setOutput('WASM not ready yet...');
      return;
    }

    try {
      if (!(window as any).run) {
        setOutput('Error: Function run() not found');
        return;
      }

      const raw = (window as any).run(input);

      if (!raw) {
        setOutput('Error: No result returned from WASM');
        return;
      }

      const result = JSON.parse(raw);

      if (
        result.parser &&
        result.parser.errors &&
        result.parser.errors.length > 0
      ) {
        setOutput(`Parser Errors:\n${result.parser.errors.join('\n')}`);
      } else if (result.output) {
        setOutput(result.output);
      } else {
        setOutput('No output generated');
      }

      (window as any).lexTokens = result.lexer || [];
      (window as any).astTree = result.ast || {};
      (window as any).parserInfo = result.parser || {};
    } catch (error) {
      setOutput(
        `Error: ${
          error instanceof Error ? error.message : 'Unknown error occurred'
        }`
      );
      console.error('Execution error:', error);
    }
  };

  return (
    <div className="flex h-screen bg-[#111] text-gray-200">
      <div className="w-1/2 p-4 border-r border-gray-700 flex flex-col">
        <div className="space-x-3.5 pb-4">
          <button>Hello world!</button>
          <button>Fibonacci</button>
          <button>Factorial</button>
          <button>Addition</button>
          <button>Map</button> <button>Hashmap</button>
          <button>Array indexing</button>{' '}
        </div>
        <div className="flex space-x-2 mb-3">
          {['clear', 'share', 'run', 'auto-run'].map((tab) => (
            <button
              key={tab}
              onClick={() => setInput('')}
              className={`px-3 py-1 rounded ${
                view === tab
                  ? 'bg-gray-200 text-black'
                  : 'bg-black border border-gray-600 text-gray-400'
              }`}
            >
              {tab.toUpperCase()}
            </button>
          ))}
        </div>
        <textarea
          value={input}
          onChange={(e) => setInput(e.target.value)}
          className="flex-1 bg-black text-gray-100 p-3 rounded-md font-mono resize-none"
          placeholder="Écris ton code ici..."
        />
        <button
          onClick={handleRun}
          className="mt-3 bg-gray-200 text-black px-4 py-2 rounded hover:bg-white transition"
        >
          Run
        </button>
      </div>

      <div className="w-1/2 p-4 flex flex-col">
        <div className="flex space-x-2 mb-3">
          {['output', 'lexer', 'ast', 'parser'].map((tab) => (
            <button
              key={tab}
              onClick={() => setView(tab)}
              className={`px-3 py-1 rounded ${
                view === tab
                  ? 'bg-gray-200 text-black'
                  : 'bg-black border border-gray-600 text-gray-400'
              }`}
            >
              {tab.toUpperCase()}
            </button>
          ))}
        </div>

        <div className="flex-1 bg-black p-3 rounded-md overflow-auto font-mono">
          {view === 'output' && <pre>{output}</pre>}
          {view === 'lexer' && (
            <pre>
              {JSON.stringify((window as any).window.lexTokens || [], null, 2)}
            </pre>
          )}
          {view === 'ast' && (
            <pre>
              {JSON.stringify((window as any).window.astTree || {}, null, 2)}
            </pre>
          )}
          {view === 'parser' && (
            <pre>
              {JSON.stringify((window as any).window.parserInfo || {}, null, 2)}
            </pre>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
