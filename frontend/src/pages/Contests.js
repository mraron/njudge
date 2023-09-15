import React, {useEffect} from 'react';
import ProfileSideBar from '../components/ProfileSidebar'
import ContestList from '../components/ContestList'
import '../index.css';

function Contests() {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar 
                        src="https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg" 
                        username="dbence"
                        score="2550"/>
                </div>
                <div className="w-full px-4 lg:pl-3">
                    <ContestList contestData={[
                        ["Online oktató programozóverseny #3", "2023-12-23, 14:00", true], 
                        ["Online oktató programozóverseny #2", "2023-08-23, 14:00", false], 
                        ["Online oktató programozóverseny #1", "2023-04-23, 14:00", false]]}/>
                </div>
            </div>
        </div>
    )
}

export default Contests;