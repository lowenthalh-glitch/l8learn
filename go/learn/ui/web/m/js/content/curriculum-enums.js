(function() {
    'use strict';
    window.MobileContentCurriculum = window.MobileContentCurriculum || {};

    MobileContentCurriculum.enums = {
        SUBJECT: { 0: 'Unknown', 1: 'Math', 2: 'Reading', 3: 'Science', 4: 'Writing', 5: 'Social Studies' },
        GRADE: { 0: 'Unknown', 1: 'Pre-K', 2: 'K', 3: '1st', 4: '2nd', 5: '3rd', 6: '4th', 7: '5th', 8: '6th', 9: '7th', 10: '8th' },
        DIFFICULTY: { 0: 'Unknown', 1: 'Intro', 2: 'Easy', 3: 'Medium', 4: 'Hard', 5: 'Challenge' },
        STATUS: { 0: 'Unknown', 1: 'Draft', 2: 'Review', 3: 'Published', 4: 'Archived' },
        STATUS_CLASSES: { 1: 'layer8d-status-pending', 2: 'layer8d-status-pending', 3: 'layer8d-status-active', 4: 'layer8d-status-inactive' },
        ACTIVITY_TYPE: { 0: 'Unknown', 1: 'Interactive', 2: 'Multiple Choice', 3: 'Free Response', 4: 'Matching', 5: 'Ordering', 6: 'Fill Blank', 7: 'Reading', 8: 'Video', 9: 'Game' }
    };

    var enums = MobileContentCurriculum.enums;

    MobileContentCurriculum.render = {
        subject: function(v) { return enums.SUBJECT[v] || v; },
        grade: function(v) { return enums.GRADE[v] || v; },
        difficulty: function(v) { return enums.DIFFICULTY[v] || v; },
        status: Layer8MRenderers.createStatusRenderer(enums.STATUS, enums.STATUS_CLASSES),
        activityType: function(v) { return enums.ACTIVITY_TYPE[v] || v; }
    };

    MobileContentCurriculum.primaryKeys = {
        Course: 'courseId',
        Unit: 'unitId',
        Lesson: 'lessonId',
        Activity: 'activityId',
        Worksheet: 'worksheetId'
    };
})();
