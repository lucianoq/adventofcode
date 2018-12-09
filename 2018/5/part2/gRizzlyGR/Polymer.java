import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Polymer {
	public static byte[] react(byte[] input) {
		int i = 0;
		int j = i + 1;

		int vainLoops = 0;
		while (true) {

			if (i >= input.length || j >= input.length) {
				i = 0;
				j = i + 1;
			}

			byte first = input[i];
			byte second = input[j];

			if (Math.abs(first - second) == 32) {
				byte[] sx = Arrays.copyOfRange(input, 0, i);
				byte[] dx = Arrays.copyOfRange(input, j + 1, input.length);
				input = merge(sx, dx);
				vainLoops = 0;
			} else {
				vainLoops++;

				// Stop if I'm looping endlessly with no more opposites
				if (vainLoops == input.length)
					break;
			}

			i = j;
			j = i + 1;

		}

		return input;
	}

	public static byte[] merge(byte[]... arrays) {
		int finalLength = 0;
		for (byte[] array : arrays) {
			finalLength += array.length;
		}

		byte[] dest = null;
		int destPos = 0;

		for (byte[] array : arrays) {
			if (dest == null) {
				dest = Arrays.copyOf(array, finalLength);
				destPos = array.length;
			} else {
				System.arraycopy(array, 0, dest, destPos, array.length);
				destPos += array.length;
			}
		}
		return dest;
	}

	public static byte[] streamline(byte pivot, byte[] bulk) {
		List<Byte> tmp = new ArrayList<>();

		for (byte b : bulk)
			tmp.add(b);

		tmp.removeAll(Arrays.asList(pivot, (byte) (pivot - 32)));

		byte[] toReturn = new byte[tmp.size()];

		for (int i = 0; i < tmp.size(); i++)
			toReturn[i] = tmp.get(i);

		return toReturn;

	}

}
