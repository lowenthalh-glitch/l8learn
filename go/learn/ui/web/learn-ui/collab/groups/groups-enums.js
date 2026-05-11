(function() {
    'use strict';
    window.CollabGroups = window.CollabGroups || {};

    CollabGroups.enums = {
        GROUP_TYPE: { 0: 'Unknown', 1: 'Study', 2: 'Project', 3: 'Challenge', 4: 'Tutoring', 5: 'Book Club', 6: 'Pod' },
        GROUP_STATUS: { 0: 'Unknown', 1: 'Active', 2: 'Completed', 3: 'Paused', 4: 'Archived' },
        GROUP_STATUS_CLASSES: { 1: 'layer8d-status-active', 2: 'layer8d-status-completed', 3: 'layer8d-status-pending', 4: 'layer8d-status-inactive' },
        MEMBER_ROLE: { 0: 'Unknown', 1: 'Leader', 2: 'Member', 3: 'Tutor', 4: 'Learner' },
        SUBJECT: { 0: 'Unknown', 1: 'Math', 2: 'Reading', 3: 'Science', 4: 'Writing', 5: 'Social Studies' },
        CONTENT_STATUS: { 0: 'Unknown', 1: 'Draft', 2: 'Review', 3: 'Published', 4: 'Archived' },
        CONTENT_STATUS_CLASSES: { 1: 'layer8d-status-pending', 2: 'layer8d-status-pending', 3: 'layer8d-status-active', 4: 'layer8d-status-inactive' },
        MASTERY_LEVEL: { 0: 'Unknown', 1: 'Not Started', 2: 'Emerging', 3: 'Developing', 4: 'Proficient', 5: 'Mastered', 6: 'Exemplary' }
    };

    var enums = CollabGroups.enums;

    CollabGroups.render = {
        groupType: function(v) { return enums.GROUP_TYPE[v] || v; },
        groupStatus: Layer8DRenderers.createStatusRenderer(enums.GROUP_STATUS, enums.GROUP_STATUS_CLASSES),
        memberRole: function(v) { return enums.MEMBER_ROLE[v] || v; },
        subject: function(v) { return enums.SUBJECT[v] || v; },
        contentStatus: Layer8DRenderers.createStatusRenderer(enums.CONTENT_STATUS, enums.CONTENT_STATUS_CLASSES)
    };

    CollabGroups.primaryKeys = {
        CollabGroup: 'groupId',
        TutorMatch: 'matchId',
        Challenge: 'challengeId'
    };
})();
