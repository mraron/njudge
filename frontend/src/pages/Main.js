import React from 'react';
import ProfileSideBar from '../components/ProfileSidebar'
import PostFrame from '../components/PostFrame'
import { SVGSpinner } from '../svg/SVGs';
import '../index.css';

function Main() {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <SVGSpinner cls="hidden w-16 h-16 absolute top-1/2 left-1/2"/>
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar 
                        src="https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg" 
                        username="dbence"
                        score="2550"/>
                </div>
                <div className="w-full px-4 lg:pl-3">
                    <PostFrame newsData={[
                        ["IOI-CEOI válogató 2023", "2023-12-23, 14:00", "Felkerültek az oldalra a 2023-as IOI-CEOI válogató feladatai."], 
                        ["Nemes Tihamér + OKTV 2023", "2023-08-23, 14:00", "Felkerültek az oldalra a 2022/2023-as Nemes Tihamér és OKTV 2. és 3. fordulóinak feladatai."], 
                        ["Kódkupa 2023 2. forduló", "2023-04-23, 14:00", "Felkerültek az oldalra a 2022/2023-as Kódkupa válogatóverseny 2. online fordulójának feladatai."]]}/>
                </div>
            </div>
        </div>
    );
}

export default Main;