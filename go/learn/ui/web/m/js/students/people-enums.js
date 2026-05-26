(function() {
    'use strict';
    window.MobilePeopleStudents = window.MobilePeopleStudents || {};

    MobilePeopleStudents.enums = {
        STUDENT_STATUS: {
            0: 'Unknown', 1: 'Active', 2: 'Inactive', 3: 'Graduated', 4: 'Transferred'
        },
        STUDENT_STATUS_VALUES: {
            'active': 1, 'inactive': 2, 'graduated': 3, 'transferred': 4
        },
        STUDENT_STATUS_CLASSES: {
            1: 'layer8d-status-active', 2: 'layer8d-status-inactive',
            3: 'layer8d-status-completed', 4: 'layer8d-status-pending'
        },
        GRADE: {
            0: 'Unknown', 1: 'Pre-K', 2: 'Kindergarten', 3: 'Grade 1', 4: 'Grade 2',
            5: 'Grade 3', 6: 'Grade 4', 7: 'Grade 5', 8: 'Grade 6', 9: 'Grade 7', 10: 'Grade 8'
        },
        GRADE_VALUES: {
            'pre-k': 1, 'kindergarten': 2, 'grade 1': 3, 'grade 2': 4, 'grade 3': 5,
            'grade 4': 6, 'grade 5': 7, 'grade 6': 8, 'grade 7': 9, 'grade 8': 10
        },
        RELATION: {
            0: 'Unknown', 1: 'Parent', 2: 'Grandparent', 3: 'Sibling', 4: 'Legal Guardian', 5: 'Other'
        },
        RELATION_VALUES: {
            'parent': 1, 'grandparent': 2, 'sibling': 3, 'legal guardian': 4, 'other': 5
        },
        TEACHER_ROLE: {
            0: 'Unknown', 1: 'Primary', 2: 'Specialist', 3: 'Aide', 4: 'Substitute', 5: 'Admin'
        },
        TEACHER_ROLE_VALUES: {
            'primary': 1, 'specialist': 2, 'aide': 3, 'substitute': 4, 'admin': 5
        },
        TEACHER_ROLE_CLASSES: {
            1: 'layer8d-status-active', 2: 'layer8d-status-pending',
            3: 'layer8d-status-pending', 4: 'layer8d-status-inactive',
            5: 'layer8d-status-completed'
        },
        EVAL_DOC_TYPE: {
            0: 'Unknown', 1: 'Speech', 2: 'Occupational Therapy', 3: 'Psychological',
            4: 'IEP', 5: 'Developmental', 6: 'Reading Specialist', 7: 'Behavioral',
            8: 'Medical', 9: 'Other'
        },
        EVAL_FINDING_STATUS: {
            0: 'Unknown', 1: 'Pending', 2: 'Accepted', 3: 'Rejected', 4: 'Edited'
        },
        EVAL_FINDING_CLASSES: {
            1: 'layer8d-status-pending', 2: 'layer8d-status-active',
            3: 'layer8d-status-inactive', 4: 'layer8d-status-pending'
        },
        EVAL_PROCESSING_STATUS: {
            0: 'Unknown', 1: 'Pending', 2: 'Extracting', 3: 'Complete', 4: 'Failed',
            5: 'Cleaned', 6: 'Submitted'
        },
        EVAL_PROCESSING_CLASSES: {
            1: 'layer8d-status-pending', 2: 'layer8d-status-pending',
            3: 'layer8d-status-active', 4: 'layer8d-status-inactive',
            5: 'layer8d-status-completed', 6: 'layer8d-status-pending'
        }
    };

    var enums = MobilePeopleStudents.enums;
    var renderEnum = Layer8MRenderers.renderEnum;

    MobilePeopleStudents.render = {
        studentStatus: Layer8MRenderers.createStatusRenderer(enums.STUDENT_STATUS, enums.STUDENT_STATUS_CLASSES),
        grade: function(v) { return renderEnum(v, enums.GRADE); },
        relation: function(v) { return renderEnum(v, enums.RELATION); },
        teacherRole: Layer8MRenderers.createStatusRenderer(enums.TEACHER_ROLE, enums.TEACHER_ROLE_CLASSES),
        evalDocType: function(v) { return enums.EVAL_DOC_TYPE[v] || v; },
        evalFindingStatus: Layer8MRenderers.createStatusRenderer(enums.EVAL_FINDING_STATUS, enums.EVAL_FINDING_CLASSES),
        evalProcessingStatus: Layer8MRenderers.createStatusRenderer(enums.EVAL_PROCESSING_STATUS, enums.EVAL_PROCESSING_CLASSES)
    };

    MobilePeopleStudents.primaryKeys = {
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
