import { Routes, Route } from "react-router-dom";

import MainLayout from "../layouts/MainLayout/MainLayout";

import Dashboard from "../pages/Dashboard/Dashboard";
import Workspace from "../pages/Workspace/Workspace";
import History from "../pages/History/History";
import Settings from "../pages/Settings/Settings";

function AppRouter() {
  return (
    <Routes>
      <Route path="/" element={<MainLayout />}>
        <Route index element={<Dashboard />} />
        <Route path="workspace" element={<Workspace />} />
        <Route path="history" element={<History />} />
        <Route path="settings" element={<Settings />} />
      </Route>
    </Routes>
  );
}

export default AppRouter;