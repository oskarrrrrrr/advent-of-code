defmodule Aoc do
  def read_input() do
    {:ok, input} = File.read("input.txt")
    String.split(input, "\n")
      |> Enum.map(fn s -> String.split(s, "   ") end)
      |> Enum.filter(fn parts -> length(parts) == 2 end)
      |> Enum.map(&List.to_tuple/1)
      |> Enum.map(fn { l, r } -> { String.to_integer(l), String.to_integer(r) } end)
      |> Enum.unzip()
  end

  def part1() do
    { left, right } = Aoc.read_input()
    Enum.zip(Enum.sort(left), Enum.sort(right))
      |> Enum.map(fn { l, r } -> abs(l-r) end)
      |> Enum.sum()
  end

  def part2() do
    { left, right } = read_input()
    right_fq = Enum.frequencies(right)
    Enum.frequencies(left)
      |> Enum.map(fn { k, v } -> k * v * (right_fq[k] || 0) end)
      |> Enum.sum()
  end
end

IO.inspect(Aoc.part1())
IO.inspect(Aoc.part2())
