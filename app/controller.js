(async function() {
    // setTimeout(() => {
    //   log("Timeout called")
    // }, 1000);
    log(Object.keys(console).join(', '));

  if (!jsProgramPath || !hash) {
    log('Usage: controller.js <module-path> <hash>');
    return;
  }

  try {
    log('Running program: ' + jsProgramPath);
    if (typeof hello === 'function') {
      hello()
    } else {
      log('The specified program does not export a default function.');
    }

    log('Hash received: ' + hash);
  } catch (err) {
    log('Error running the program:', err.message);
  }
})();
