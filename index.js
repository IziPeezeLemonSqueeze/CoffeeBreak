const RealIdle = require('@paymoapp/real-idle');
const robot = require('robotjs');

const args = process.argv.slice(2);

setInterval(() => {
    const idleStatus = RealIdle.getIdleState(300);
    const idleTime = RealIdle.getIdleSeconds();

    if (args[0] == '--d')
    {
        console.log('System idle state:', idleStatus);
        console.log('  - Idle seconds:', idleTime);
    }


    if (idleStatus === 'idle' || idleTime >= 290)
    {
        robot.keyTap('shift');
    }
}, 5000); 