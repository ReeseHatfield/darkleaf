import { app, BrowserWindow } from 'electron';
import * as path from 'path';

function createWindow() {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'),
      nodeIntegration: true,
      contextIsolation: false,
    },
  });

  win.loadURL(`file://${path.join(__dirname, 'index.html')}`);
}

app.whenReady().then(createWindow);

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    // do some extra stuff here to ensure save and encrpyt
    app.quit();
  }
});

app.on('activate', () => {
  if (BrowserWindow.getAllWindows().length === 0) {
    // ensure key and stuff
    createWindow();
  }
});
