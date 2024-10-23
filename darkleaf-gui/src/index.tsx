import React, { useEffect, useState } from 'react';
import ReactDOM from 'react-dom';
import { ipcRenderer } from 'electron';

const App = () => {
  const [args, setArgs] = useState<string[]>([]);

  useEffect(() => {
    ipcRenderer.on('send-args', (event, flags) => {
      setArgs(flags);
    });
  }, []);

  return (
    <div>
      <h1>Command-Line Flags</h1>
      <p>Passed flags: {args.join(', ')}</p>
    </div>
  );
};

ReactDOM.render(<App />, document.getElementById('root'));
