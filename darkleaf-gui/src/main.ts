import { app, BrowserWindow } from 'electron';
import * as path from 'path';

function createWindow() {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      nodeIntegration: true, 
      contextIsolation: false, 
    },
  });

  win.setMenu(null)

  win.loadFile(path.join(__dirname, 'index.html'));

  // fancy ipc message passing
  win.webContents.on('did-finish-load', () => {
    const flags = process.argv.slice(2) // skip node and electron args
    win.webContents.send('send-args', flags)
  })
}

app.whenReady().then(createWindow);

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

app.on('activate', () => {
  if (BrowserWindow.getAllWindows().length === 0) {
    createWindow();
  }
});
