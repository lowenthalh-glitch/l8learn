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
        ]),
        StudentProfile: f.form('Student Profile', [
            f.section('Summary', [
                ...f.reference('studentId', 'Student', 'Student'),
                ...f.textarea('overallDescription', 'Overall Description'),
                ...f.text('mainStrengths', 'Main Strengths'),
                ...f.text('mainChallenges', 'Main Challenges'),
                ...f.text('primaryGoals', 'Primary Goals')
            ]),
            f.section('Readiness Scores', [
                ...f.number('readiness.academicReadiness', 'Academic Readiness'),
                ...f.number('readiness.readingReadiness', 'Reading Readiness'),
                ...f.number('readiness.mathReadiness', 'Math Readiness'),
                ...f.number('readiness.writingFineMotor', 'Writing/Fine Motor'),
                ...f.number('readiness.speechLanguage', 'Speech/Language'),
                ...f.number('readiness.attentionStamina', 'Attention Stamina'),
                ...f.number('readiness.socialEmotional', 'Social-Emotional'),
                ...f.number('readiness.independence', 'Independence')
            ]),
            f.section('Learning Style', [
                ...f.text('learningStyle.preferredModes', 'Preferred Modes'),
                ...f.number('learningStyle.bestSessionLengthMinutes', 'Best Session (min)'),
                ...f.number('learningStyle.bestActivityLengthMinutes', 'Best Activity (min)'),
                ...f.number('learningStyle.breakFrequencyMinutes', 'Break Frequency (min)'),
                ...f.text('learningStyle.bestTimeOfDay', 'Best Time of Day')
            ]),
            f.section('Attention', [
                ...f.number('attention.focusPreferredActivityMinutes', 'Focus Preferred Activity (min)'),
                ...f.number('attention.focusAcademicTaskMinutes', 'Focus Academic Task (min)'),
                ...f.text('attention.losingFocusSigns', 'Losing Focus Signs'),
                ...f.text('attention.helpfulSupports', 'Helpful Supports')
            ]),
            f.section('Math', [
                ...f.text('math.level', 'Level'),
                ...f.text('math.addition', 'Addition'),
                ...f.text('math.subtraction', 'Subtraction'),
                ...f.text('math.multiplication', 'Multiplication'),
                ...f.text('math.division', 'Division'),
                ...f.text('math.fractions', 'Fractions'),
                ...f.text('math.errorPatterns', 'Error Patterns')
            ]),
            f.section('Literacy', [
                ...f.text('literacy.readingLevel', 'Reading Level'),
                ...f.text('literacy.letterRecognition', 'Letter Recognition'),
                ...f.text('literacy.phonemicAwareness', 'Phonemic Awareness'),
                ...f.text('literacy.comprehension', 'Comprehension'),
                ...f.number('literacy.readingFluencyWpm', 'Fluency (WPM)')
            ]),
            f.section('AI Tutor Settings', [
                ...f.text('aiTutor.personality', 'Personality'),
                ...f.text('aiTutor.shouldDo', 'Should Do'),
                ...f.text('aiTutor.shouldAvoid', 'Should Avoid'),
                ...f.text('aiTutor.sentenceLength', 'Sentence Length'),
                ...f.text('aiTutor.hintStyle', 'Hint Style')
            ])
        ]),
        EvalImport: f.form('Evaluation Import', [
            f.section('Evaluation Info', [
                ...f.reference('studentId', 'Student', 'Student'),
                ...f.select('documentType', 'Document Type', enums.EVAL_DOC_TYPE),
                ...f.text('professionalName', 'Professional Name'),
                ...f.date('evaluationDate', 'Evaluation Date'),
                ...f.checkbox('allReviewed', 'All Reviewed'),
                ...f.number('acceptedCount', 'Accepted'),
                ...f.number('rejectedCount', 'Rejected'),
                ...f.checkbox('appliedToProfile', 'Applied to Profile')
            ])
        ])
    };
})();
