import ProfileSideBar from '../components/ProfileSidebar'
import DropdownListFrame from '../components/DropdownListFrame'

function Archive() {
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
                    <div className="mb-3">
                        <DropdownListFrame title="Nemes Tihamér 1. kategória" tree={{
                            "children":
                            [
                                {
                                    "title": "Nemes Tihamér 1. kategória",
                                    "children": 
                                    [
                                        {
                                            "title": "2020 2. forduló",
                                            "children": [
                                                {"title": "Kövek (100)"},
                                                {"title": "Micimackó (40)"},
                                                {"title": "Morze (60)"},
                                                {"title": "Völgy (100)"}
                                            ]
                                        },
                                        {
                                            "title": "2020 3. forduló",
                                            "children": [
                                                {"title": "Kövek (100)"},
                                                {"title": "Micimackó (40)"},
                                                {"title": "Morze (60)"},
                                                {"title": "Völgy (100)"}
                                            ]
                                        }
                                    ]
                                }
                            ]}}/>
                    </div>
                    <div className="mb-3">
                        <DropdownListFrame title="OKTV" tree={{
                            "children":
                            [
                                {
                                    "title": "2020",
                                    "children": 
                                    [
                                        {
                                            "title": "2020 2. forduló",
                                            "children": [
                                                {"title": "Bürokrácia (40)"},
                                                {"title": "Gyros (30)"},
                                                {"title": "JardaT"},
                                                {"title": "Jegesmedve (50)"},
                                                {"title": "Múzeumi őrök"},
                                                {"title": "Tom és Jerry2 (60)"},
                                                {"title": "Utazás (40)"}
                                            ]
                                        }
                                    ]
                                }
                            ]}}/>
                        </div>
                        <div className="mb-3">
                            <DropdownListFrame title="Nemes Tihamér 1. kategória" tree={{
                                "children":
                                [
                                    {
                                        "title": "Nemes Tihamér 1. kategória",
                                        "children": 
                                        [
                                            {
                                                "title": "2020 2. forduló",
                                                "children": [
                                                    {"title": "Kövek (100)"},
                                                    {"title": "Micimackó (40)"},
                                                    {"title": "Morze (60)"},
                                                    {"title": "Völgy (100)"}
                                                ]
                                            },
                                            {
                                                "title": "2020 3. forduló",
                                                "children": [
                                                    {"title": "Kövek (100)"},
                                                    {"title": "Micimackó (40)"},
                                                    {"title": "Morze (60)"},
                                                    {"title": "Völgy (100)"}
                                                ]
                                            }
                                        ]
                                    }
                                ]}}/>
                        </div>
                </div>
            </div>
        </div>
    );
}

export default Archive;