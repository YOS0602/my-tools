import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import Container from "@mui/material/Container";
import App from "./App.tsx";
import "./index.css";
import "@fontsource/noto-sans-jp/400.css";

// MUI Component で使用するデフォルトフォントを設定
const theme = createTheme({
  typography: {
    fontFamily: ["Noto Sans JP", "sans-serif"].join(","),
  },
});

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ThemeProvider theme={theme}>
      <Container maxWidth="lg">
        <App />
      </Container>
    </ThemeProvider>
  </StrictMode>
);
