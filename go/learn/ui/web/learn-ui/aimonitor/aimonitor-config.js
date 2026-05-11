(function() {
    'use strict';
    var svc = Layer8ModuleConfigFactory.service;
    var mod = Layer8ModuleConfigFactory.module;

    Layer8ModuleConfigFactory.create({
        namespace: 'AIMonitor',
        modules: {
            'data': mod('Data', '\uD83E\uDD16', [
                { key: 'promptlogs', label: 'Prompt Logs', icon: '\uD83D\uDCCB', endpoint: '/30/PromptLog', model: 'LLMPromptLog', readOnly: true },
                svc('llmconfig', 'LLM Config', '\u2699\uFE0F', '/30/LLMConfig', 'LLMConfig')
            ])
        },
        submodules: ['AIMonitorData']
    });
})();
