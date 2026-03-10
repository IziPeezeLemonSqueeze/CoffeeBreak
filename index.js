const { execSync, spawnSync } = require('child_process');
const fs = require('fs');
const path = require('path');
const os = require('os');

const args = process.argv.slice(2);
const debug = args[0] === '--d';

// Scrivi lo script PS1 su disco una volta sola
const psScript = path.join(os.tmpdir(), 'get_idle.ps1');
fs.writeFileSync(psScript, `
Add-Type -MemberDefinition @'
    [DllImport("user32.dll")]
    public static extern bool GetLastInputInfo(ref LASTINPUTINFO plii);
    [StructLayout(LayoutKind.Sequential)]
    public struct LASTINPUTINFO {
        public uint cbSize;
        public uint dwTime;
    }
'@ -Name 'User32' -Namespace 'Win32'

$lii = New-Object Win32.User32+LASTINPUTINFO
$lii.cbSize = [System.Runtime.InteropServices.Marshal]::SizeOf($lii)
[Win32.User32]::GetLastInputInfo([ref]$lii) | Out-Null
$idle = [Environment]::TickCount - $lii.dwTime
[math]::Floor($idle / 1000)
`);

function getIdleSeconds() {
    const result = spawnSync('powershell', [
        '-ExecutionPolicy', 'Bypass',
        '-File', psScript
    ], { windowsHide: true });
    return parseInt(result.stdout.toString().trim()) || 0;
}

function pressShift() {
    spawnSync('powershell', [
        '-ExecutionPolicy', 'Bypass',
        '-Command', "Add-Type -AssemblyName System.Windows.Forms; [System.Windows.Forms.SendKeys]::SendWait('+')"
    ], { windowsHide: true });
}

setInterval(() => {
    const idleTime = getIdleSeconds();
    const idleStatus = idleTime >= 300 ? 'idle' : 'active';

    if (debug) {
        console.log('System idle state:', idleStatus);
        console.log('  - Idle seconds:', idleTime);
    }

    if (idleTime >= 290) {
        if (debug) console.log('  → Premo Shift...');
        pressShift();
    }
}, 5000);