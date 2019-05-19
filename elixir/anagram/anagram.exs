defmodule Anagram do
  @doc """
  Returns all candidates that are anagrams of, but not equal to, 'base'.
  """
  @spec match(String.t(), [String.t()]) :: [String.t()]
  def match(base, candidates) do
    Enum.filter(candidates, &anagrams?(&1, base))
  end

  defp anagrams?(word1, word2) do
    String.downcase(word1) != String.downcase(word2) &&
      sort(word1) == sort(word2)
  end

  defp sort(word) do
    String.downcase(word)
    |> String.graphemes
    |> Enum.sort
  end
end
