import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class Day2Part2 {

	public static void main(String[] args) throws IOException {
		String inputPath = "";

		if (args.length != 0)
			inputPath = args[0];
		else
			inputPath = "input";

		List<String> ids = Files.readAllLines(Paths.get(inputPath));

		for (int i = 0; i < ids.size(); i++) {
			char[] first = ids.get(i).toCharArray();

			for (int j = 0; j < ids.size(); j++) {

				if (i == j)
					continue;

				char[] second = ids.get(j).toCharArray();

				findCommonElements(first, second);
			}
		}

	}

	public static void findCommonElements(char[] first, char[] second) {

		if (first == null || second == null)
			throw new IllegalArgumentException("Arguments must not be null");

		if (first.length != second.length)
			throw new IllegalArgumentException("Input arrays must have the same number of elements");

		int numDifferentElems = 0;
		List<Character> commonElems = new ArrayList<>();

		for (int i = 0; i < first.length; i++) {
			if (first[i] == second[i])
				commonElems.add(first[i]);
			else
				numDifferentElems++;
		}

		if (numDifferentElems == 1)
			System.out.println(
					String.join("", commonElems
							.stream()
							.map(Object::toString)
							.collect(Collectors.joining(""))));

	}
}
