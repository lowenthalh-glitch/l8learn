/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
*/
(function() {
    'use strict';
    window.StudentsPeople = window.StudentsPeople || {};

    var STUDENT_STATUS = {
        0: 'Unknown',
        1: 'Active',
        2: 'Inactive',
        3: 'Graduated',
        4: 'Transferred'
    };

    var STUDENT_STATUS_VALUES = {
        'active': 1,
        'inactive': 2,
        'graduated': 3,
        'transferred': 4
    };

    var STUDENT_STATUS_CLASSES = {
        1: 'layer8d-status-active',
        2: 'layer8d-status-inactive',
        3: 'layer8d-status-completed',
        4: 'layer8d-status-pending'
    };

    var GRADE = {
        0: 'Unknown',
        1: 'Pre-K',
        2: 'Kindergarten',
        3: 'Grade 1',
        4: 'Grade 2',
        5: 'Grade 3',
        6: 'Grade 4',
        7: 'Grade 5',
        8: 'Grade 6',
        9: 'Grade 7',
        10: 'Grade 8'
    };

    var GRADE_VALUES = {
        'pre-k': 1,
        'kindergarten': 2,
        'grade 1': 3,
        'grade 2': 4,
        'grade 3': 5,
        'grade 4': 6,
        'grade 5': 7,
        'grade 6': 8,
        'grade 7': 9,
        'grade 8': 10
    };

    var RELATION = {
        0: 'Unknown',
        1: 'Parent',
        2: 'Grandparent',
        3: 'Sibling',
        4: 'Legal Guardian',
        5: 'Other'
    };

    var RELATION_VALUES = {
        'parent': 1,
        'grandparent': 2,
        'sibling': 3,
        'legal guardian': 4,
        'other': 5
    };

    var TEACHER_ROLE = {
        0: 'Unknown',
        1: 'Primary',
        2: 'Specialist',
        3: 'Aide',
        4: 'Substitute',
        5: 'Admin'
    };

    var TEACHER_ROLE_VALUES = {
        'primary': 1,
        'specialist': 2,
        'aide': 3,
        'substitute': 4,
        'admin': 5
    };

    var TEACHER_ROLE_CLASSES = {
        1: 'layer8d-status-active',
        2: 'layer8d-status-pending',
        3: 'layer8d-status-pending',
        4: 'layer8d-status-inactive',
        5: 'layer8d-status-completed'
    };

    StudentsPeople.enums = {
        STUDENT_STATUS: STUDENT_STATUS,
        STUDENT_STATUS_VALUES: STUDENT_STATUS_VALUES,
        STUDENT_STATUS_CLASSES: STUDENT_STATUS_CLASSES,
        GRADE: GRADE,
        GRADE_VALUES: GRADE_VALUES,
        RELATION: RELATION,
        RELATION_VALUES: RELATION_VALUES,
        TEACHER_ROLE: TEACHER_ROLE,
        TEACHER_ROLE_VALUES: TEACHER_ROLE_VALUES,
        TEACHER_ROLE_CLASSES: TEACHER_ROLE_CLASSES
    };

    var renderEnum = Layer8DRenderers.renderEnum;
    var createStatusRenderer = Layer8DRenderers.createStatusRenderer;

    StudentsPeople.render = {
        studentStatus: createStatusRenderer(STUDENT_STATUS, STUDENT_STATUS_CLASSES),
        grade: function(value) { return renderEnum(value, GRADE); },
        relation: function(value) { return renderEnum(value, RELATION); },
        teacherRole: createStatusRenderer(TEACHER_ROLE, TEACHER_ROLE_CLASSES)
    };

    var EVAL_DOC_TYPE = {
        0: 'Unknown', 1: 'Speech', 2: 'Occupational Therapy', 3: 'Psychological',
        4: 'IEP', 5: 'Developmental', 6: 'Reading Specialist', 7: 'Behavioral',
        8: 'Medical', 9: 'Other'
    };

    var EVAL_FINDING_STATUS = {
        0: 'Unknown', 1: 'Pending', 2: 'Accepted', 3: 'Rejected', 4: 'Edited'
    };

    var EVAL_FINDING_CLASSES = {
        1: 'layer8d-status-pending', 2: 'layer8d-status-active',
        3: 'layer8d-status-inactive', 4: 'layer8d-status-pending'
    };

    StudentsPeople.enums.EVAL_DOC_TYPE = EVAL_DOC_TYPE;
    StudentsPeople.enums.EVAL_FINDING_STATUS = EVAL_FINDING_STATUS;
    StudentsPeople.enums.EVAL_FINDING_CLASSES = EVAL_FINDING_CLASSES;

    StudentsPeople.render.evalDocType = function(v) { return EVAL_DOC_TYPE[v] || v; };
    StudentsPeople.render.evalFindingStatus = Layer8DRenderers.createStatusRenderer(EVAL_FINDING_STATUS, EVAL_FINDING_CLASSES);

    StudentsPeople.primaryKeys = {
        Student: 'studentId',
        StudentProfile: 'profileId',
        Guardian: 'guardianId',
        Teacher: 'teacherId',
        Classroom: 'classroomId',
        School: 'schoolId',
        District: 'districtId',
        EvalImport: 'importId'
    };
})();
