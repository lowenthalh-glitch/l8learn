(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var render = MobileDataAIMonitor.render;

    MobileDataAIMonitor.columns = {
        LLMPromptLog: [
            ...col.id('logId'),
            ...col.enum('type', 'Type', null, render.promptType),
            ...col.col('studentId', 'Student'),
            ...col.boolean('containsPii', 'Contains PII'),
            ...col.number('systemPromptTokens', 'Sys Tokens'),
            ...col.number('userMessageTokens', 'User Tokens'),
            ...col.number('responseTimeMs', 'Response (ms)'),
            ...col.date('timestamp', 'Timestamp')
        ],
        LLMConfig: [
            ...col.id('configId'),
            ...col.status('mode', 'Mode', null, render.mode),
            ...col.col('apiProvider', 'Provider'),
            ...col.col('modelName', 'Model'),
            ...col.boolean('piiMaskingEnabled', 'PII Masking'),
            ...col.number('maxDailyCalls', 'Max Daily'),
            ...col.number('callsToday', 'Calls Today')
        ]
    };

    MobileDataAIMonitor.columns.LLMPromptLog[1].primary = true;
    MobileDataAIMonitor.columns.LLMPromptLog[3].secondary = true;

    MobileDataAIMonitor.columns.LLMConfig[1].primary = true;
    MobileDataAIMonitor.columns.LLMConfig[2].secondary = true;
})();
