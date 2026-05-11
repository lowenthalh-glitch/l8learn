(function() {
    'use strict';
    Layer8SectionConfigs.register('aimonitor', {
        title: 'AI Monitor',
        subtitle: 'LLM Prompts & Safety',
        icon: '\uD83E\uDD16',
        initFn: 'initializeAIMonitor',
        modules: [{
            key: 'data',
            label: 'Data',
            icon: '\uD83E\uDD16',
            isDefault: true,
            services: [
                { key: 'promptlogs', label: 'Prompt Logs', icon: '\uD83D\uDCCB', isDefault: true },
                { key: 'llmconfig', label: 'LLM Config', icon: '\u2699\uFE0F' }
            ]
        }]
    });
})();
