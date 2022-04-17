package config

import (
	"flag"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

const (
	Reset = "\033[0m"

	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"

	BackBlack  = "\033[40m"
	BackRed    = "\033[41m"
	BackGreen  = "\033[42m"
	BackYellow = "\033[43m"
	BackBlue   = "\033[44m"
	BackPurple = "\033[45m"
	BackCyan   = "\033[46m"
	BackGray   = "\033[47m"
	BackWhite  = "\033[107m"
)

func GenerateFlags() {
	/* true */ FILTRAR_VOLUME_FINANC := flag.Bool("volfin", true, "Usado para filtrar Volume Financeiro de cada ação.")
	/* true */ FILTRAR_MARGEM_EBIT := flag.Bool("mebit", true, "Usado para filtrar Margem Ebit de cada ação.")
	/* true */ FILTRAR_PL := flag.Bool("pl", true, "Usado para filtrar Pl de cada ação.")
	/* true */ FILTRAR_ROA := flag.Bool("roa", true, "Usado para filtrar Roa de cada ação.")

	/* 200.000 */
	VOL_FIN_MIN := flag.Int("vl-min", 200000, "Volume financeiro minimo, uma ação com volume financeiro baixo é pouco negociada na bolsa, e pode gerar dificuldades na hora da venda")
	/* 0 */ MARGEM_EBIT_MINIMA := flag.Float64("mebit-min", 0.0, "Margem Ebit mininima, uma ação com margem ebit negativa diz que a empresa dá prejuízo, queremos retirá-las")
	/* 1.5 */ PL_MINIMO := flag.Float64("pl-min", 1.5, "Pl minimo, quanto menor, melhor, entretanto valores pequenos podem indicar que os dados nao estao bons, geralmente utiliza-se valores entre 1 e 2")
	/* 5 */ ROA_MINIMO := flag.Float64("roa-min", 5.0, "Roa minimo, medido em porcentagem, considera-se um roa alto, bom")

	/* 1.5 */
	PESO_PL := flag.Float64("p-pl", 1.5, "Peso do Pl para rankear as ações")
	/* 1 */ PESO_ROA := flag.Float64("p-roa", 1.0, "Peso do Roa para rankear as ações")
	/* 2 */ PESO_EV_EBIT := flag.Float64("p-evebit", 2.0, "Peso do EvEbit para rankear as ações")

	/* cli */
	OUTPUT := flag.String("out", "cli", "Fomato de arquivo a ser gerado na saida. Suporte para txt, xlsx, cli.")
	flag.Parse()

	viper.Set("FILTRAR_VOLUME_FINANC", *FILTRAR_VOLUME_FINANC)
	viper.Set("FILTRAR_MARGEM_EBIT", *FILTRAR_MARGEM_EBIT)
	viper.Set("FILTRAR_PL", *FILTRAR_PL)
	viper.Set("FILTRAR_ROA", *FILTRAR_ROA)
	viper.Set("VOL_FIN_MIN", *VOL_FIN_MIN)
	viper.Set("MARGEM_EBIT_MINIMA", *MARGEM_EBIT_MINIMA)
	viper.Set("PL_MINIMO", *PL_MINIMO)
	viper.Set("ROA_MINIMO", *ROA_MINIMO)
	viper.Set("PESO_PL", *PESO_PL)
	viper.Set("PESO_ROA", *PESO_ROA)
	viper.Set("PESO_EV_EBIT", *PESO_EV_EBIT)
	viper.Set("OUTPUT", *OUTPUT)
}

func HttpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}

	return client
}
