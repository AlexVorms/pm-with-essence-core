import { LayoutDashboard, Users, Settings, Plus } from 'lucide-react';

import {BoardPage} from "./BoardPage.tsx";

export const HomePage = () => {
    return (
        <div className="flex h-screen bg-white text-gray-700 border-r border-gray-200">
            {/* Вертикальный Navbar */}
            <nav className="w-16 flex flex-col items-center py-4 space-y-8 bg-slate-50">
                <div className="p-2 bg-blue-600 rounded-lg text-white">
                    <LayoutDashboard size={24} />
                </div>

                <div className="flex flex-col space-y-4">
                    <NavItem icon={<Users size={20} />} />
                    <NavItem icon={<Settings size={20} />} />
                    <NavItem icon={<Plus size={20} />} />
                </div>
            </nav>

            {/* Область контента (здесь будет ваша доска) */}
            <main className="flex-1 overflow-x-auto bg-[#f1f2f4]">
                <div className="p-4">
                    <BoardPage boardId={"aaaaaaaa-1111-1111-1111-aaaaaaaaaaaa"} projectId={"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"}></BoardPage>
                </div>
            </main>
        </div>
    );
};

// @ts-ignore
const NavItem = ({ icon }) => (
    <button className="p-2 hover:bg-gray-200 rounded-md transition-colors">
        {icon}
    </button>
);