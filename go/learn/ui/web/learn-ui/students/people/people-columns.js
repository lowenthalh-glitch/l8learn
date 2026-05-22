/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
*/
(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var enums = StudentsPeople.enums;
    var render = StudentsPeople.render;

    StudentsPeople.columns = {
        Student: [
            ...col.id('studentId'),
            ...col.col('firstName', 'First Name'),
            ...col.col('lastName', 'Last Name'),
            ...col.enum('gradeLevel', 'Grade', enums.GRADE_VALUES, render.grade),
            ...col.status('status', 'Status', enums.STUDENT_STATUS_VALUES, render.studentStatus),
            ...col.col('classroomId', 'Classroom'),
            ...col.col('schoolId', 'School')
        ],
        Guardian: [
            ...col.id('guardianId'),
            ...col.col('firstName', 'First Name'),
            ...col.col('lastName', 'Last Name'),
            ...col.col('email', 'Email'),
            ...col.enum('relation', 'Relation', enums.RELATION_VALUES, render.relation)
        ],
        Teacher: [
            ...col.id('teacherId'),
            ...col.col('firstName', 'First Name'),
            ...col.col('lastName', 'Last Name'),
            ...col.col('email', 'Email'),
            ...col.status('role', 'Role', enums.TEACHER_ROLE_VALUES, render.teacherRole),
            ...col.col('schoolId', 'School')
        ],
        Classroom: [
            ...col.id('classroomId'),
            ...col.col('name', 'Name'),
            ...col.enum('gradeLevel', 'Grade', enums.GRADE_VALUES, render.grade),
            ...col.col('primaryTeacherId', 'Primary Teacher'),
            ...col.col('schoolId', 'School')
        ],
        School: [
            ...col.id('schoolId'),
            ...col.col('name', 'Name'),
            ...col.col('districtId', 'District')
        ],
        District: [
            ...col.id('districtId'),
            ...col.col('name', 'Name'),
            ...col.col('stateProvince', 'State/Province')
        ],
        StudentProfile: [
            ...col.id('profileId'),
            ...col.col('studentId', 'Student'),
            ...col.col('shortSummary', 'Summary'),
            ...col.number('scores.overallAcademicReadiness', 'Academic'),
            ...col.number('scores.readingReadiness', 'Reading'),
            ...col.number('scores.mathReadiness', 'Math'),
            ...col.date('lastUpdated', 'Last Updated')
        ],
        EvalImport: [
            ...col.id('importId'),
            ...col.col('studentId', 'Student'),
            ...col.enum('documentType', 'Type', null, render.evalDocType),
            ...col.col('professionalName', 'Professional'),
            ...col.date('evaluationDate', 'Eval Date'),
            ...col.status('processingStatus', 'Status', null, render.evalProcessingStatus),
            ...col.custom('filePath', 'Document', function(item) {
                if (!item.filePath) return '-';
                return '<a href="#" onclick="event.stopPropagation();LearnFileViewer.view(\'' +
                    item.filePath.replace(/'/g, "\\'") + '\',\'Cleaned Evaluation\');return false;" ' +
                    'style="color:var(--layer8d-primary);text-decoration:underline;">View</a>';
            }),
            ...col.number('acceptedCount', 'Accepted'),
            ...col.number('rejectedCount', 'Rejected')
        ]
    };
})();
