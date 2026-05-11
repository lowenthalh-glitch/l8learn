(function() {
    'use strict';
    window.AIMonitorData = window.AIMonitorData || {};

    AIMonitorData.enums = {
        PROMPT_TYPE: { 0: 'Unknown', 1: 'Path Decision', 2: 'Profile Update', 3: 'Risk Assessment', 4: 'Progress Summary', 5: 'Parent Coaching', 6: 'Worksheet Scan', 7: 'Content Analysis', 8: 'Schedule', 9: 'Chat', 10: 'Moderation', 11: 'Eval Import' },
        LLM_MODE: { 0: 'Unknown', 1: 'Live', 2: 'Simulate', 3: 'Log Only' },
        MODE_CLASSES: { 1: 'layer8d-status-active', 2: 'layer8d-status-pending', 3: 'layer8d-status-inactive' }
    };

    var enums = AIMonitorData.enums;

    AIMonitorData.render = {
        promptType: function(v) { return enums.PROMPT_TYPE[v] || v; },
        mode: Layer8DRenderers.createStatusRenderer(enums.LLM_MODE, enums.MODE_CLASSES)
    };

    AIMonitorData.primaryKeys = {
        LLMPromptLog: 'logId',
        LLMConfig: 'configId'
    };
})();
