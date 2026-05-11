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

    StudentsPeople.primaryKeys = {
        Student: 'studentId',
        Guardian: 'guardianId',
        Teacher: 'teacherId',
        Classroom: 'classroomId',
        School: 'schoolId',
        District: 'districtId'
    };
})();
