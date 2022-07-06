# VSA Processor
The Verify, Send and Archive processor is a service that checks to see if the incoming emails are being processed, and if not, check to see if they are new clients. If they are new clients, or mismatched VSA workflows, attempt to reconfigure to properly sort incoming emails

## Main Logic
- If email is unprocessed and has a pre-exisiting VSA workflow, assume we weren't supposed to catch this email and do nothing
- If email is unprocessed and does not have a pre-existing VSA workflow, but does have a matching inbox based on the `From address domain`, then add a VSA workflow to point incoming emails to the matching inbox
- If email is unprocessed and does not have a VSA workflow, and does not have a matching inbox, create an inbox and a VSA workflow to point at the new inbox

All processes will create logs on workflow and inbox changes, as well as alert engineers to their status incase there needs to be a logical change or special case consideration. 