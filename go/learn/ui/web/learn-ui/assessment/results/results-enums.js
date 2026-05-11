(function() {
    'use strict';
    window.AssessmentResults = window.AssessmentResults || {};

    AssessmentResults.enums = {
        SESSION_STATUS: { 0: 'Unknown', 1: 'Active', 2: 'Completed', 3: 'Abandoned' },
        SESSION_STATUS_CLASSES: { 1: 'layer8d-status-active', 2: 'layer8d-status-completed', 3: 'layer8d-status-inactive' },
        INTERACTION_RESULT: { 0: 'Unknown', 1: 'Correct', 2: 'Incorrect', 3: 'Partial', 4: 'Skipped', 5: 'Hint Used', 6: 'Timed Out' },
        BENCHMARK_TYPE: { 0: 'Unknown', 1: 'Diagnostic', 2: 'Formative', 3: 'Summative', 4: 'Progress Monitor' },
        SCAN_STATUS: { 0: 'Unknown', 1: 'Uploaded', 2: 'Processing', 3: 'Review', 4: 'Complete', 5: 'Failed' },
        SCAN_STATUS_CLASSES: { 1: 'layer8d-status-pending', 2: 'layer8d-status-pending', 3: 'layer8d-status-pending', 4: 'layer8d-status-active', 5: 'layer8d-status-terminated' },
        CONTENT_STATUS: { 0: 'Unknown', 1: 'Draft', 2: 'Review', 3: 'Published', 4: 'Archived' },
        CONTENT_STATUS_CLASSES: { 1: 'layer8d-status-pending', 2: 'layer8d-status-pending', 3: 'layer8d-status-active', 4: 'layer8d-status-inactive' }
    };

    var enums = AssessmentResults.enums;

    AssessmentResults.render = {
        sessionStatus: Layer8DRenderers.createStatusRenderer(enums.SESSION_STATUS, enums.SESSION_STATUS_CLASSES),
        interactionResult: function(v) { return enums.INTERACTION_RESULT[v] || v; },
        benchmarkType: function(v) { return enums.BENCHMARK_TYPE[v] || v; },
        scanStatus: Layer8DRenderers.createStatusRenderer(enums.SCAN_STATUS, enums.SCAN_STATUS_CLASSES),
        contentStatus: Layer8DRenderers.createStatusRenderer(enums.CONTENT_STATUS, enums.CONTENT_STATUS_CLASSES)
    };

    AssessmentResults.primaryKeys = {
        LearningSession: 'sessionId',
        Score: 'scoreId',
        Benchmark: 'benchmarkId',
        WorksheetScan: 'scanId'
    };
})();
