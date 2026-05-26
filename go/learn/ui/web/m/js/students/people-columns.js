(function() {
    'use strict';
    var col = window.Layer8ColumnFactory;
    var enums = MobilePeopleStudents.enums;
    var render = MobilePeopleStudents.render;

    MobilePeopleStudents.columns = {
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
            ...col.number('acceptedCount', 'Accepted'),
            ...col.number('rejectedCount', 'Rejected')
        ]
    };

    // Primary and secondary markers for mobile card display
    MobilePeopleStudents.columns.Student[1].primary = true;
    MobilePeopleStudents.columns.Student[4].secondary = true;
    MobilePeopleStudents.columns.Guardian[1].primary = true;
    MobilePeopleStudents.columns.Guardian[4].secondary = true;
    MobilePeopleStudents.columns.Teacher[1].primary = true;
    MobilePeopleStudents.columns.Teacher[4].secondary = true;
    MobilePeopleStudents.columns.Classroom[1].primary = true;
    MobilePeopleStudents.columns.Classroom[2].secondary = true;
    MobilePeopleStudents.columns.School[1].primary = true;
    MobilePeopleStudents.columns.School[2].secondary = true;
    MobilePeopleStudents.columns.District[1].primary = true;
    MobilePeopleStudents.columns.District[2].secondary = true;
    MobilePeopleStudents.columns.StudentProfile[1].primary = true;
    MobilePeopleStudents.columns.StudentProfile[2].secondary = true;
    MobilePeopleStudents.columns.EvalImport[1].primary = true;
    MobilePeopleStudents.columns.EvalImport[5].secondary = true;
})();
