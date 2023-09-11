import SubmissionsTable from "../../components/SubmissionsTable";
import Checkbox from "../../components/Checkbox"
import RoundedFrame from "../../components/RoundedFrame";
import Pagination from "../../components/Pagination";

function SubmissionFilterFrame() {
    return (
        <RoundedFrame>
            <div className="px-6 py-4 flex flex-col sm:flex-row items-start sm:items-center justify-between">
                <div className="mb-2 sm:mb-0">
                    <Checkbox label="Teljes megoldások" />
                </div>
                <Checkbox label="Saját beküldéseim" />
            </div>
        </RoundedFrame>
    )
}

function ProblemSubmissions() {
    const submissions = [
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Áruszállítás üres szakaszai", "cpp17", "Hibás válasz 2/50", "381 ms", "21320 KiB"],
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Tom és Jerry2 (60)", "cpp17", "Hibás válasz 50/50", "381 ms", "21320 KiB"],
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Kaktuszgráf", "cpp17", "Hibás válasz 60/60", "381 ms", "21320 KiB"],
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Kaktuszgráf", "cpp17", "Hibás válasz 2/50", "381 ms", "21320 KiB"],
    ];
    return (
        <div>
            <div className="mb-3">
                <SubmissionFilterFrame />
            </div>
            <div className="mb-2">
                <SubmissionsTable submissions={submissions} />
            </div>
            <Pagination current={1000} last={2000} />
        </div>
    )
}

export default ProblemSubmissions;