import 'styled-components'

// Tipando as cores para que o TypeScrip reconheça
declare module 'styled-components'{
    export interface DefaultTheme {
        title: string;
        colors: {
            primary: string;
            secondary: string;
            tertiary: string;
            blackOne: string;
            blacktwo: string;
            blackthree: string;
        };
    }
}