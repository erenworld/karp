import { type Dispatch, type SetStateAction } from 'react';
import { Editor } from '@monaco-editor/react';

const CodeInputContent = ({
  input,
  setInput,
}: {
  input?: string;
  setInput: Dispatch<SetStateAction<string | undefined>>;
}) => (
  <Editor
    defaultLanguage="karp"
    options={{
      minimap: {
        enabled: false,
      },
    }}
    value={input}
    onChange={(value) => setInput(value)}
  />
);

export default CodeInputContent;
