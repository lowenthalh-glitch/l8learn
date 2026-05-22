/*
© 2025 Sharon Aicler (saichler@gmail.com)
Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.

Desktop finding review popup — thin wrapper around L8EvalReview core.
*/
(function() {
    'use strict';

    window.L8EvalReviewDesktop = {

        open: function(evalImport, serviceConfig) {
            var content = L8EvalReview.renderFindingsHtml(evalImport);

            Layer8DPopup.show({
                title: 'Review Evaluation Findings',
                content: content,
                size: 'xlarge',
                showFooter: evalImport.processingStatus === 3,
                saveButtonText: 'Apply Approved to Profile',
                onSave: function() {
                    var payload = L8EvalReview.buildApplyPayload(evalImport);
                    var endpoint = Layer8DConfig.resolveEndpoint(serviceConfig.endpoint);
                    fetch(endpoint, {
                        method: 'PUT',
                        headers: { 'Content-Type': 'application/json', 'Authorization': 'Bearer ' + sessionStorage.getItem('bearerToken') },
                        body: JSON.stringify(payload)
                    }).then(function(resp) {
                        if (resp.ok) {
                            Layer8DNotification.success('Findings applied to profile');
                            Layer8DPopup.close();
                        } else {
                            Layer8DNotification.error('Failed to apply findings');
                        }
                    });
                },
                onShow: function(body) {
                    body.addEventListener('click', function(e) {
                        var btn = e.target.closest('[data-action]');
                        if (!btn) return;
                        var idx = parseInt(btn.getAttribute('data-idx'));
                        var action = btn.getAttribute('data-action');
                        L8EvalReview.onFindingAction(evalImport, idx, action);
                        body.innerHTML = L8EvalReview.renderFindingsHtml(evalImport);
                    });
                }
            });
        }
    };
})();
