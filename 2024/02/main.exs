defmodule Aoc do
  def read_input() do
    {:ok, input} = File.read("input.txt")
    String.split(input, "\n", trim: true)
      |> Enum.map(fn s -> String.split(s, " ") end)
      |> Enum.map(fn l -> Enum.map(l, &String.to_integer/1) end)
  end

  defp _is_safe(report, cond) do
    case report do
      [] -> true
      [_] -> true
      [prev, next | tail] -> cond.(prev, next) and _is_safe([next | tail], cond)
    end
  end

  defp is_safe(report) do
    if _is_safe(report, &</2) or _is_safe(report, &>/2) do
      _is_safe(report, fn prev, next -> abs(prev - next) <= 3 end)
    else
      false
    end
  end

  def part1() do
    Aoc.read_input()
    |> Enum.map(&is_safe/1)
    |> Enum.count(& &1)
  end

  defp _is_safe_dampened(report, prev_prev, cond, dampening) do
    case report do
      [] -> true
      [_] -> true
      [prev, next | tail] ->
        if cond.(prev, next) do
          _is_safe_dampened([next | tail], prev, cond, dampening)
        else
          dampening > 0 and (
            _is_safe_dampened([prev | tail], prev_prev, cond, dampening - 1)
            or (
              (prev_prev == nil or cond.(prev_prev, next))
              and _is_safe_dampened([next | tail], nil, cond, dampening - 1)
            )
          )
        end
    end
  end

  defp is_safe_dampened(report) do
    inc = _is_safe_dampened(report, nil, fn prev, next -> prev < next and abs(prev - next) <= 3 end, 1)
    dec = _is_safe_dampened(report, nil, fn prev, next -> prev > next and abs(prev - next) <= 3 end, 1)
    inc or dec
  end

  def part2() do
    Aoc.read_input()
    |> Enum.map(&is_safe_dampened/1)
    |> Enum.count(& &1)
  end
end

IO.inspect(Aoc.part1())
IO.inspect(Aoc.part2())
