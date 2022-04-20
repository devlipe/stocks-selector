package config

var helpMessage map[string]string

//Needs a test that gives a bool, and base on the result return a value for true(t) and false(f)
func Ternary[K any](test bool, t K, f K) K {
	if test {
		return t
	} else {
		return f
	}
}

func init() {
	helpMessage = make(map[string]string)
	helpMessage["1"] = `TRADING VOLUME
	* Trading volume is the total number of shares of a security traded during a given period of time.
	* Trading volume can provide investors with a signal to enter the market.
	* Trading volume can also signal when an investor should take profits and sell a security due to low activity.

	In this case, we use it to indicate that a stock have a good activity and it will not be a problem when trying to sell it. So, it its good practice to filter it

	For more info: https://en.wikipedia.org/wiki/Volume_(finance)
	`
	helpMessage["2"] = `Margin EBIT
	These financial metrics measure levels and rates of profitability. 
	Probably the most common way to determine the successfulness of a company is to look at the net profits of the business. 
	
	Net profit: To calculate net profit for a unit (such as a company or division), subtract all costs, including a fair share of total corporate overheads, from the gross revenues.

			Net profit ($) = Sales revenue ($) - Total costs ($)

	We use this filter to select the companies that are profitable, this is, the ones that have a positive Margin Ebit

	For more info: https://en.wikipedia.org/wiki/Operating_margin
	`
	helpMessage["3"] = `ROA
	This number tells you what the company can do with what it has, i.e. how many dollars of earnings they derive from each dollar of assets they control. 
	It's a useful number for comparing competing companies in the same industry.
	ROA can be computed as below:

			ROA = Net Income / Average Total Assets

	This is used to select stocks that can make profit form their assets.

	For more info: https://en.wikipedia.org/wiki/Return_on_assets
	`
	helpMessage["4"] = `Price-earnings ratio
	The ratio is used for valuing companies and to find out whether they are overvalued or undervalued.
	As an example, if share A is trading at $24 and the earnings per share for the most recent 12-month period is $3, then share A has a P/E ratio of $24/($3 per year) = 8.
	Put another way, the purchaser of the share is investing $8 for every dollar of annual earnings; or, if earnings stayed constant it would take 8 years to recoup the share price.
	PE can be computed as below:

			PE = Share Price / Earnings per share

	This is used to select stocks that have a low price compared to the profit they make. Tha is the cheaper stocks.

	For more info: https://en.wikipedia.org/wiki/Price–earnings_ratio
	`
	helpMessage["5"] = helpMessage["1"]
	helpMessage["6"] = helpMessage["2"]
	helpMessage["7"] = helpMessage["3"]
	helpMessage["8"] = helpMessage["4"]

	helpMessage["9"] = `EV/Ebit Weight
	Use this to determine the importance of EV/Ebit indicator.

	A greater number will make the shares with a lower Ev/Ebit show upper in the table
	`
	helpMessage["10"] = `ROA Weight
	Use this to determine the importance of ROA indicator.

	A greater number will make the shares with a higher ROA show upper in the table
	`
	helpMessage["11"] = `PE Weight
	Use this to determine the importance of PE indicator.

	A greater number will make the shares with a lower PE show upper in the table
	`
	helpMessage["12"] = `Output
	Use this to determine the output of the program. We currently support 
		* cli (wich is shown on terminal)
		* txt (create a txt file on the current directory)
		* xlsx (create a xlsx file on the current directory)
	`
}
