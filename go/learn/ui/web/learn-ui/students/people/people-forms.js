/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
*/
(function() {
    'use strict';
    var f = window.Layer8FormFactory;
    var enums = StudentsPeople.enums;

    StudentsPeople.forms = {
        Student: f.form('Student', [
            f.section('Student Information', [
                ...f.text('firstName', 'First Name', true),
                ...f.text('lastName', 'Last Name', true),
                ...f.select('gradeLevel', 'Grade Level', enums.GRADE, true),
                ...f.select('status', 'Status', enums.STUDENT_STATUS),
                ...f.reference('classroomId', 'Classroom', 'Classroom'),
                ...f.reference('schoolId', 'School', 'School'),
                ...f.reference('districtId', 'District', 'District')
            ])
        ]),
        Guardian: f.form('Guardian', [
            f.section('Guardian Information', [
                ...f.text('firstName', 'First Name', true),
                ...f.text('lastName', 'Last Name', true),
                ...f.text('email', 'Email'),
                ...f.text('phone', 'Phone'),
                ...f.select('relation', 'Relation', enums.RELATION)
            ])
        ]),
        Teacher: f.form('Teacher', [
            f.section('Teacher Information', [
                ...f.text('firstName', 'First Name', true),
                ...f.text('lastName', 'Last Name', true),
                ...f.text('email', 'Email'),
                ...f.select('role', 'Role', enums.TEACHER_ROLE),
                ...f.reference('schoolId', 'School', 'School')
            ])
        ]),
        Classroom: f.form('Classroom', [
            f.section('Classroom Information', [
                ...f.text('name', 'Name', true),
                ...f.select('gradeLevel', 'Grade Level', enums.GRADE),
                ...f.reference('primaryTeacherId', 'Primary Teacher', 'Teacher'),
                ...f.reference('schoolId', 'School', 'School'),
                ...f.text('academicYear', 'Academic Year')
            ])
        ]),
        School: f.form('School', [
            f.section('School Information', [
                ...f.text('name', 'Name', true),
                ...f.reference('districtId', 'District', 'District')
            ])
        ]),
        District: f.form('District', [
            f.section('District Information', [
                ...f.text('name', 'Name', true),
                ...f.text('stateProvince', 'State/Province')
            ])
        ])
    };
})();
