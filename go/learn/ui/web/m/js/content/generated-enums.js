(function() {
    'use strict';
    window.MobileContentGenerated = window.MobileContentGenerated || {};

    var LESSON_STATUS = {
        0: 'Unknown', 1: 'Generating', 2: 'Ready', 3: 'In Progress', 4: 'Completed', 5: 'Skipped'
    };
    var LESSON_STATUS_CLASSES = {
        1: 'layer8d-status-pending', 2: 'layer8d-status-active',
        3: 'layer8d-status-pending', 4: 'layer8d-status-completed',
        5: 'layer8d-status-inactive'
    };

    MobileContentGenerated.enums = {
        LESSON_STATUS: LESSON_STATUS,
        LESSON_STATUS_CLASSES: LESSON_STATUS_CLASSES
    };

    MobileContentGenerated.render = {
        lessonStatus: Layer8MRenderers.createStatusRenderer(LESSON_STATUS, LESSON_STATUS_CLASSES)
    };

    MobileContentGenerated.primaryKeys = {
        GeneratedLesson: 'generatedLessonId'
    };
})();
