import Pagination from "../../components/Pagination";
import SubmissionsTable from "../../components/SubmissionsTable";

function ProfileSubmissions() {
    const submissions = [
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Áruszállítás üres szakaszai", "cpp17", "Hibás válasz 2/50", "381 ms", "21320 KiB"],
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Tom és Jerry2 (60)", "cpp17", "Hibás válasz 50/50", "381 ms", "21320 KiB"],
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Kaktuszgráf", "cpp17", "Hibás válasz 60/60", "381 ms", "21320 KiB"],
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Kaktuszgráf", "cpp17", "Hibás válasz 2/50", "381 ms", "21320 KiB"],
    ];
    return (
        <div>
            <div className="mb-2">
                <SubmissionsTable submissions={submissions} />
            </div>
            <Pagination current={1000} last={2000} />
        </div>
    );
}

export default ProfileSubmissions;