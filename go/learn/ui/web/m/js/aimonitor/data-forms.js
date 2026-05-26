(function() {
    'use strict';
    var f = window.Layer8FormFactory;
    var enums = MobileDataAIMonitor.enums;

    MobileDataAIMonitor.forms = {
        LLMPromptLog: f.form('LLM Prompt Log', [
            f.section('Prompt Info', [
                ...f.select('type', 'Type', enums.PROMPT_TYPE, false, { readOnly: true }),
                ...f.text('studentId', 'Student', false, { readOnly: true }),
                ...f.select('mode', 'Mode', enums.LLM_MODE, false, { readOnly: true }),
                ...f.textarea('systemPrompt', 'System Prompt', { readOnly: true }),
                ...f.textarea('userMessage', 'User Message', { readOnly: true }),
                ...f.textarea('response', 'Response', { readOnly: true }),
                ...f.number('systemPromptTokens', 'System Prompt Tokens', false, { readOnly: true }),
                ...f.number('userMessageTokens', 'User Message Tokens', false, { readOnly: true }),
                ...f.number('responseTokens', 'Response Tokens', false, { readOnly: true }),
                ...f.number('responseTimeMs', 'Response Time (ms)', false, { readOnly: true }),
                ...f.checkbox('containsStudentName', 'Contains Student Name', { readOnly: true }),
                ...f.checkbox('containsPii', 'Contains PII', { readOnly: true }),
                ...f.checkbox('dataMasked', 'Data Masked', { readOnly: true })
            ])
        ]),
        LLMConfig: f.form('LLM Config', [
            f.section('Configuration', [
                ...f.select('mode', 'Mode', enums.LLM_MODE, true),
                ...f.text('apiProvider', 'API Provider', true),
                ...f.text('modelName', 'Model Name', true),
                ...f.number('maxTokens', 'Max Tokens'),
                ...f.number('temperature', 'Temperature'),
                ...f.checkbox('piiMaskingEnabled', 'PII Masking Enabled'),
                ...f.checkbox('promptLoggingEnabled', 'Prompt Logging Enabled'),
                ...f.number('maxDailyCalls', 'Max Daily Calls'),
                ...f.number('callsToday', 'Calls Today', false, { readOnly: true })
            ])
        ])
    };
})();
