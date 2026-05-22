/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Shared batch-process logic — platform-independent.
Used by desktop (people-batch-process.js) and mobile (people-batch-process-m.js).
*/
(function() {
    'use strict';

    window.L8BatchProcess = {

        // Upload multiple files, creating one EvalImport per file.
        // uploadFn(file) must return a Promise resolving to { storagePath }.
        // postFn(evalData) must return a Promise resolving on success.
        // onProgress(current, total, fileName) called after each file.
        uploadFiles: async function(files, studentId, uploadFn, postFn, onProgress) {
            var results = { success: 0, failed: 0, errors: [] };
            for (var i = 0; i < files.length; i++) {
                var file = files[i];
                if (onProgress) onProgress(i + 1, files.length, file.name);
                try {
                    var uploadResult = await uploadFn(file);
                    await postFn({
                        studentId: studentId,
                        filePath: uploadResult.storagePath,
                        documentType: 0
                    });
                    results.success++;
                } catch (err) {
                    results.failed++;
                    results.errors.push(file.name + ': ' + err.message);
                }
            }
            return results;
        },

        // Trigger batch LLM processing by PUTting one CLEANED eval with SUBMITTED status.
        // evals: array of eval objects from the table data.
        // putFn(eval) must return a Promise.
        triggerProcessAll: async function(evals, putFn) {
            var cleaned = evals.filter(function(e) { return e.processingStatus === 5; }); // CLEANED=5
            if (cleaned.length === 0) {
                return { triggered: false, message: 'No cleaned evaluations ready to process' };
            }
            // PUT only the first CLEANED eval with SUBMITTED status
            var trigger = cleaned[0];
            trigger.processingStatus = 6; // SUBMITTED=6
            try {
                await putFn(trigger);
                return { triggered: true, message: 'Processing ' + cleaned.length + ' evaluations...' };
            } catch (err) {
                return { triggered: false, message: 'Failed to trigger processing: ' + err.message };
            }
        },

        // Get counts of evals by status for a student.
        getStatusCounts: function(evals) {
            var counts = { pending: 0, extracting: 0, complete: 0, failed: 0, cleaned: 0, submitted: 0 };
            for (var i = 0; i < evals.length; i++) {
                var s = evals[i].processingStatus;
                if (s === 1) counts.pending++;
                else if (s === 2) counts.extracting++;
                else if (s === 3) counts.complete++;
                else if (s === 4) counts.failed++;
                else if (s === 5) counts.cleaned++;
                else if (s === 6) counts.submitted++;
            }
            return counts;
        },

        // Build multi-file upload HTML (no platform-specific elements).
        renderUploadFormHtml: function(studentRefHtml) {
            return '<div class="l8-batch-upload">' +
                '<div class="form-group">' + studentRefHtml + '</div>' +
                '<div class="form-group">' +
                '<label>Evaluation Documents (PDF)</label>' +
                '<div class="l8-file-drop-area" id="l8-batch-dropzone" ' +
                'style="border:2px dashed var(--layer8d-border);border-radius:8px;padding:24px;text-align:center;cursor:pointer;">' +
                '<div style="color:var(--layer8d-text-medium);margin-bottom:8px;">Drop PDF files here or click to browse</div>' +
                '<div style="font-size:12px;color:var(--layer8d-text-muted);">Multiple files supported (max 5MB each)</div>' +
                '<input type="file" id="l8-batch-files" multiple accept=".pdf" style="display:none;">' +
                '</div>' +
                '<div id="l8-batch-file-list" style="margin-top:8px;"></div>' +
                '</div>' +
                '</div>';
        },

        // Render file list preview.
        renderFileList: function(files) {
            if (!files || files.length === 0) return '';
            var html = '<div style="font-size:13px;">';
            for (var i = 0; i < files.length; i++) {
                var size = files[i].size > 1024 ? Math.round(files[i].size / 1024) + ' KB' : files[i].size + ' B';
                html += '<div style="padding:4px 0;border-bottom:1px solid var(--layer8d-border);">' +
                    Layer8DUtils.escapeHtml(files[i].name) + ' <span style="color:var(--layer8d-text-muted);">(' + size + ')</span></div>';
            }
            html += '</div>';
            return html;
        }
    };
})();
