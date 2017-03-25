defmodule Bob do
  def hey(input) do
    cond do
      silence?(input) ->
        "Fine. Be that way!"
      question?(input) ->
        "Sure."
      shouting?(input) ->
        "Whoa, chill out!"
      true ->
        "Whatever."
    end
  end

  defp silence?(input) do
    String.strip(input) == ""
  end

  defp shouting?(input) do
    String.upcase(input) == input && String.downcase(input) != input
  end

  defp question?(input) do
    String.ends_with?(input, "?")
  end
end
