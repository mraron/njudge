import Pagination from '../components/Pagination';
import ProfileSideBar from '../components/ProfileSidebar'
import SubmissionsTable from '../components/SubmissionsTable';

function Submissions() {
    const submissions = [
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Áruszállítás üres szakaszai", "cpp17", "Hibás válasz 2/50", "381 ms", "21320 KiB"],
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Tom és Jerry2 (60)", "cpp17", "Hibás válasz 50/50", "381 ms", "21320 KiB"],
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Kaktuszgráf", "cpp17", "Hibás válasz 60/60", "381 ms", "21320 KiB"],
        ["5669", "2023-09-06, 14:23:42", "gonterarmin", "Kaktuszgráf", "cpp17", "Hibás válasz 2/50", "381 ms", "21320 KiB"],
    ];
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar 
                        src="https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg" 
                        username="dbence"
                        score="2550"/>
                </div>
                <div className="w-full px-4 lg:pl-3 overflow-x-auto">
                    <div className="mb-2">
                        <SubmissionsTable submissions={submissions} />
                    </div>
                    <Pagination current={1000} last={2000} />
                </div>
            </div>
        </div>
    );
}

export default Submissions;