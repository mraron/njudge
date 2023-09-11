import React from 'react';
import { ProblemFilterFrame } from '../components/ProblemFilter';
import ProblemsTable from '../components/ProblemsTable'
import ProfileSideBar from '../components/ProfileSidebar'
import Pagination from '../components/Pagination';
import '../index.css';

function Problems() {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar 
                        src="https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg" 
                        username="dbence"
                        score="2550"/>
                </div>
                <div className="w-full flex flex-col overflow-x-auto">
                    <div className="w-full px-4 lg:pl-3">
                        <div className="mb-2">
                            <ProblemFilterFrame />
                        </div>
                        <div className="mb-2">
                            <ProblemsTable />
                        </div>
                        <Pagination current={1} last={2000} />
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Problems;